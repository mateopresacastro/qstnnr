package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/mateopresacastro/qstnnr/pkg/api"
	"github.com/mateopresacastro/qstnnr/pkg/qerr"
	"github.com/mateopresacastro/qstnnr/pkg/qservice"
	"github.com/mateopresacastro/qstnnr/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// To assert implementation
var _ api.QuestionnaireServer = (*server)(nil)

// server implements the gRPC questionnaire service.
type server struct {
	api.QuestionnaireServer
	service qservice.QService
	logger  *slog.Logger
}

// ServerConfig holds the configuration for the gRPC server.
type Config struct {
	Logger  *slog.Logger
	Service qservice.QService
}

// New creates a new gRPC server with the given configuration.
func New(cfg *Config) (*grpc.Server, error) {
	server := &server{service: cfg.Service, logger: cfg.Logger}
	grpcsrv := grpc.NewServer()
	api.RegisterQuestionnaireServer(grpcsrv, server)
	return grpcsrv, nil
}

// GetQuestions returns all questions with their options.
func (s *server) GetQuestions(ctx context.Context, _ *emptypb.Empty) (*api.GetQuestionsResponse, error) {
	qsts, err := s.service.Questions()
	if err != nil {
		return nil, s.handleError(err)
	}

	var questions []*api.Question
	for qID, q := range qsts {
		var options []*api.Option
		var question = &api.Question{Id: int32(qID), Text: q.Text, Options: options}
		for oID, o := range q.Options {
			question.Options = append(question.Options, &api.Option{Id: int32(oID), Text: o.Text})
		}
		questions = append(questions, question)
	}

	return &api.GetQuestionsResponse{Questions: questions}, nil
}

// SubmitAnswers processes submitted answers and returns results with statistics.
func (s *server) SubmitAnswers(ctx context.Context, req *api.SubmitAnswersRequest) (*api.SubmitAnswersResponse, error) {
	answers := make(map[store.QuestionID]store.OptionID)
	for _, a := range req.Answers {
		answers[store.QuestionID(a.QuestionId)] = store.OptionID(a.OptionId)
	}
	result, err := s.service.SubmitAnswers(answers)
	if err != nil {
		return nil, s.handleError(err)
	}

	processed, err := s.processSolutions(result.Solutions)
	if err != nil {
		return nil, s.handleError(err)
	}
	return &api.SubmitAnswersResponse{
		Solutions:  processed,
		BetterThan: int32(result.Stat),
		Correct:    int32(result.Correct),
	}, nil
}

// GetSolutions returns the correct answers for all questions.
func (s *server) GetSolutions(ctx context.Context, req *emptypb.Empty) (*api.GetSolutionsResponse, error) {
	solutions, err := s.service.Solutions()
	if err != nil {
		return nil, s.handleError(err)
	}

	processed, err := s.processSolutions(solutions)
	if err != nil {
		return nil, s.handleError(err)
	}

	return &api.GetSolutionsResponse{Solutions: processed}, nil
}

// processSolutions converts internal solution format to API response format.
func (s *server) processSolutions(ss map[store.QuestionID]store.OptionID) ([]*api.Solution, error) {
	qsts, err := s.service.Questions()
	if err != nil {
		return nil, err
	}
	var processed []*api.Solution
	for qID, oID := range ss {
		q := &api.Question{Id: int32(qID), Text: qsts[qID].Text}
		s := &api.Solution{
			Question: q, CorrectOptionId: int32(oID),
			CorrectOptionText: qsts[qID].Options[oID].Text,
		}
		processed = append(processed, s)
	}
	return processed, nil
}

// handleError centralizes the error handling. It maps domain level error codes
// to gRPC codes and reports bugs.
func (s *server) handleError(err error) error {
	unknownError := status.Error(codes.Unknown, "an unexpected error occurred")

	// If this is not a ServiceError we know is not a known edge case and is a real bug.
	serviceErr, ok := err.(qservice.ServiceError)
	if !ok {
		s.reportBug(err)
		return unknownError
	}

	// Get the inner error from ServiceError and check if it's a QError
	if qErr, ok := errors.Unwrap(serviceErr).(qerr.QError); ok {
		grpcCode, ok := errorCodeToGRPC[qErr.Code]
		if !ok {
			s.reportBug(fmt.Errorf("error mapping domain error code %d to gRPC error code", qErr.Code))
			return unknownError
		}
		return status.Error(grpcCode, qErr.Message)
	}

	return status.Error(codes.Internal, serviceErr.Error())
}

// reportBug logs unexpected errors for debugging. Here we could send this bug to a
// centralized destination.
func (s *server) reportBug(err error) {
	msg := "there was an unnespected issue; please report this as a bug"
	s.logger.Error(msg, "err", fmt.Sprintf("%#v", err))
}

// errorCodeToGRPC maps domain level errors to gRPC errors.
var errorCodeToGRPC = map[qerr.ErrorCode]codes.Code{
	qerr.Unknown:      codes.Unknown,
	qerr.InvalidInput: codes.InvalidArgument,
	qerr.NotFound:     codes.NotFound,
	qerr.Internal:     codes.Internal,
}
