package cases

// import (
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/zakether/go-interview-bot/internal/cases/mocks"
// 	"github.com/zakether/go-interview-bot/internal/entities"
// 	"go.uber.org/zap"
// )

// func TestQuestionService_StartTest(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStorage := mocks.NewMockSessionStorage(ctrl)
// 	mockLogger := zap.NewNop()

// 	questionService := NewQuestionService(mockStorage, mockLogger)

// 	mockSession := &entities.Session{}
// 	mockStorage.EXPECT().CreateSession(mockSession).Return(nil)

// 	session, err := questionService.StartTest()

// 	assert.NoError(t, err)
// 	assert.NotNil(t, session)
// }

// func TestQuestionService_GetResult(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStorage := mocks.NewMockSessionStorage(ctrl)
// 	mockLogger := zap.NewNop()

// 	questionService := NewQuestionService(mockStorage, mockLogger)

// 	session := &entities.Session{}
// 	userAnswers := []int{1, 2, 3}

// 	mockStorage.EXPECT().SubmitUserAnswers(session, userAnswers).Return(nil)

// 	score, err := questionService.GetResult(session, userAnswers)

// 	assert.NoError(t, err)
// 	assert.Equal(t, session.Score, score)
// }

// func TestQuestionService_GetResult_Error(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStorage := mocks.NewMockSessionStorage(ctrl)
// 	mockLogger := zap.NewNop()

// 	questionService := NewQuestionService(mockStorage, mockLogger)

// 	session := &entities.Session{}
// 	userAnswers := []int{1, 2, 3}

// 	expectedErr := errors.New("storage error")
// 	mockStorage.EXPECT().SubmitUserAnswers(session, userAnswers).Return(expectedErr)

// 	score, err := questionService.GetResult(session, userAnswers)

// 	assert.Error(t, err)
// 	assert.Zero(t, score)
// }
