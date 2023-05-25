package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	rankService *rank.RankService
}

// NewPermissionController new language controller.
func NewPermissionController(rankService *rank.RankService) *PermissionController {
	return &PermissionController{rankService: rankService}
}

// GetPermission check user permission
// @Summary check user permission
// @Description check user permission
// @Tags Permission
// @Security ApiKeyAuth
// @Param Authorization header string true "access-token"
// @Produce json
// @Param action query string true "permission key" Enums(question.add, question.edit, question.edit_without_review, question.delete, question.close, question.reopen, question.vote_up, question.vote_down, question.pin, question.unpin, question.hide, question.show, answer.add, answer.edit, answer.edit_without_review, answer.delete, answer.accept, answer.vote_up, answer.vote_down, answer.invite_someone_to_answer, comment.add, comment.edit, comment.delete, comment.vote_up, comment.vote_down, report.add, tag.add, tag.edit, tag.edit_slug_name, tag.edit_without_review, tag.delete, tag.synonym, link.url_limit, vote.detail, answer.audit, question.audit, tag.audit, tag.use_reserved_tag)
// @Success 200 {object} handler.RespBody{data=map[string]bool}
// @Router /answer/api/v1/permission [get]
func (u *PermissionController) GetPermission(ctx *gin.Context) {
	req := &schema.GetPermissionReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	userID := middleware.GetLoginUserIDFromContext(ctx)
	resp, err := u.rankService.CheckOperationPermissions(ctx, userID, req.Actions)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	mapping := make(map[string]bool, len(resp))
	for i, action := range req.Actions {
		mapping[action] = resp[i]
	}
	handler.HandleResponse(ctx, err, mapping)
}
