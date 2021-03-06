
import { ApiTypes, CommonTypes } from 'src/types'

export enum Types {
  GET_MESSAGES_REQUEST = 'GET_MESSAGES_REQUEST',
  GET_MESSAGES_SUCCESS = 'GET_MESSAGES_SUCCESS',

  GET_MORE_MESSAGES_REQUEST = 'GET_MORE_MESSAGES_REQUEST',
  GET_MORE_MESSAGES_SUCCESS = 'GET_MORE_MESSAGES_SUCCESS',
  GET_MORE_MESSAGES_FAILED = 'GET_MORE_MESSAGES_FAILED',

  GET_CURRENT_HUB_REQUEST = 'GET_CURRENT_HUB_REQUEST',
  GET_CURRENT_HUB_SUCCESS = 'GET_CURRENT_HUB_SUCCESS',
  GET_CURRENT_HUB_FAILED = 'GET_CURRENT_HUB_FAILED',

  GET_MESSAGES_FROM_HUB_REQUEST = 'GET_MESSAGES_FROM_HUB_REQUEST',
  GET_MESSAGES_FROM_HUB_SUCCESS = 'GET_MESSAGES_FROM_HUB_SUCCESS',
  GET_MESSAGES_FROM_HUB_FAILED = 'GET_MESSAGES_FROM_HUB_FAILED',
  
  GET_MORE_MESSAGES_FROM_HUB_REQUEST = 'GET_MORE_MESSAGES_FROM_HUB_REQUEST',
  GET_MORE_MESSAGES_FROM_HUB_SUCCESS = 'GET_MORE_MESSAGES_FROM_HUB_SUCCESS',
  GET_MORE_MESSAGES_FROM_HUB_FAILED = 'GET_MORE_MESSAGES_FROM_HUB_FAILED',
  
  GET_MESSAGE_BY_ID_FROM_HUB_REQUEST = 'GET_MESSAGE_BY_ID_FROM_HUB_REQUEST',
  GET_MESSAGE_BY_ID_FROM_HUB_SUCCESS = 'GET_MESSAGE_BY_ID_FROM_HUB_SUCCESS',
  GET_MESSAGE_BY_ID_FROM_HUB_FAILED = 'GET_MESSAGE_BY_ID_FROM_HUB_FAILED',

  RESET_MESSAGE_BY_ID = 'RESET_MESSAGE_BY_ID',
  
  CLEAN_ALL_MESSAGES = 'CLEAN_ALL_MESSAGES',

  POST_MESSAGE_REQUEST = 'POST_MESSAGE_REQUEST',
  POST_MESSAGE_SUCCESS = 'POST_MESSAGE_SUCCESS',

  DELETE_MESSAGE_REQUEST = 'DELETE_MESSAGE_REQUEST',
  DELETE_MESSAGE_SUCCESS = 'DELETE_MESSAGE_SUCCESS',
  
  REPORT_MESSAGE_HUB_REQUEST = 'REPORT_MESSAGE_HUB_REQUEST',
  REPORT_MESSAGE_CENTRAL_REQUEST = 'REPORT_MESSAGE_CENTRAL_REQUEST',
  REPORT_MESSAGE_SUCCESS = 'REPORT_MESSAGE_SUCCESS',

  EDIT_MESSAGE_REQUEST = 'EDIT_MESSAGE_REQUEST',
  EDIT_MESSAGE_SUCCESS = 'EDIT_MESSAGE_SUCCESS',

  POST_COMMENT_REQUEST = 'POST_COMMENT_REQUEST',
  POST_COMMENT_SUCCESS = 'POST_COMMENT_SUCCESS',

  EDIT_COMMENT_REQUEST = 'EDIT_COMMENT_REQUEST',
  EDIT_COMMENT_SUCCESS = 'EDIT_COMMENT_SUCCESS',

  DELETE_COMMENT_REQUEST = 'DELETE_COMMENT_REQUEST',
  DELETE_COMMENT_SUCCESS = 'DELETE_COMMENT_SUCCESS',

  GET_MESSAGE_UPLOAD_LINK_REQUEST = 'GET_MESSAGE_UPLOAD_LINK_REQUEST',
  GET_MESSAGE_UPLOAD_LINK_SUCCESS = 'GET_MESSAGE_UPLOAD_LINK_SUCCESS',

  SET_MESSAGE_ATTACHMENT_REQUEST = 'SET_MESSAGE_ATTACHMENT_REQUEST',
  SET_MESSAGE_ATTACHMENT_SUCCESS = 'SET_MESSAGE_ATTACHMENT_SUCCESS',

  LIKE_MESSAGE_REQUEST = 'LIKE_MESSAGE_REQUEST',
  LIKE_MESSAGE_SUCCESS = 'LIKE_MESSAGE_SUCCESS',
  
  LIKE_COMMENT_REQUEST = 'LIKE_COMMENT_REQUEST',
  LIKE_COMMENT_SUCCESS = 'LIKE_COMMENT_SUCCESS',
  
  GET_LIKES_FOR_MESSAGE_REQUEST = 'GET_LIKES_FOR_MESSAGE_REQUEST',
  GET_LIKES_FOR_MESSAGE_SUCCESS = 'GET_LIKES_FOR_MESSAGE_SUCCESS',
  
  GET_LIKES_FOR_COMMENT_REQUEST = 'GET_LIKES_FOR_COMMENT_REQUEST',
  GET_LIKES_FOR_COMMENT_SUCCESS = 'GET_LIKES_FOR_COMMENT_SUCCESS',

  HIDE_MESSAGE_REQUEST = 'HIDE_MESSAGE_REQUEST',
  HIDE_MESSAGE_SUCCESS = 'HIDE_MESSAGE_SUCCESS',

  HIDE_COMMENT_REQUEST = 'HIDE_COMMENT_REQUEST',
  HIDE_COMMENT_SUCCESS = 'HIDE_COMMENT_SUCCESS',
}

const getMessagesRequest = () => ({
  type: Types.GET_MESSAGES_REQUEST,
})

const getMessagesSuccess = (payload: CommonTypes.HubTypes.CurrentHub[]) => ({
  type: Types.GET_MESSAGES_SUCCESS,
  payload
})

const getCurrentHubRequest = () => ({
  type: Types.GET_CURRENT_HUB_REQUEST,
})

const getCurrentHubSuccess = (payload: CommonTypes.HubTypes.CurrentHub) => ({
  type: Types.GET_CURRENT_HUB_SUCCESS,
  payload
})

const getCurrentHubFailed = () => ({
  type: Types.GET_CURRENT_HUB_FAILED,
})

const getMessagesFromHubRequest = (payload: ApiTypes.Messages.MessagesFromHub) => ({
  type: Types.GET_MESSAGES_FROM_HUB_REQUEST,
  payload,
})

const getMessagesFromHubSuccess = (payload: {
  hub: string
  messages: ApiTypes.Messages.Message[]
}) => ({
  type: Types.GET_MESSAGES_FROM_HUB_SUCCESS,
  payload
})

const getMessagesFromHubFailed = () => ({
  type: Types.GET_MESSAGES_FROM_HUB_FAILED
})

const postMessageRequest = (payload: ApiTypes.Messages.PostMessage) => ({
  type: Types.POST_MESSAGE_REQUEST,
  payload
})

const postMessageSucces = (payload: boolean) => ({
  type: Types.POST_MESSAGE_SUCCESS,
  payload,
})

const deleteMessageRequest = (payload: ApiTypes.Messages.DeleteMessage) => ({
  type: Types.DELETE_MESSAGE_REQUEST,
  payload
})

const deleteMessageSucces = () => ({
  type: Types.DELETE_MESSAGE_SUCCESS,
})

const reportMessageHubRequest = (payload: ApiTypes.Messages.ReportMessageHub) => ({
  type: Types.REPORT_MESSAGE_HUB_REQUEST,
  payload
})

const reportMessageCentralRequest = (payload: ApiTypes.Messages.ReportMessageCentral) => ({
  type: Types.REPORT_MESSAGE_CENTRAL_REQUEST,
  payload
})

const reportMessageSucces = () => ({
  type: Types.REPORT_MESSAGE_SUCCESS,
})

const editMessageRequest = (payload: ApiTypes.Messages.EditMessage) => ({
  type: Types.EDIT_MESSAGE_REQUEST,
  payload
})

const editMessageSucces = () => ({
  type: Types.EDIT_MESSAGE_SUCCESS,
})

const postCommentRequest = (payload: ApiTypes.Messages.PostComment) => ({
  type: Types.POST_COMMENT_REQUEST,
  payload
})

const postCommentSucces = (payload: boolean) => ({
  type: Types.POST_COMMENT_SUCCESS,
  payload,
})

const editCommentRequest = (payload: ApiTypes.Messages.EditComment) => ({
  type: Types.EDIT_COMMENT_REQUEST,
  payload
})

const editCommentSucces = () => ({
  type: Types.EDIT_COMMENT_SUCCESS,
})

const deleteCommentRequest = (payload: ApiTypes.Messages.DeleteComment) => ({
  type: Types.DELETE_COMMENT_REQUEST,
  payload
})

const deleteCommentSucces = () => ({
  type: Types.DELETE_COMMENT_SUCCESS,
})

const getMessageUploadLinkRequest = (payload: ApiTypes.Messages.UploadLinkRequest) => ({
  type: Types.GET_MESSAGE_UPLOAD_LINK_REQUEST,
  payload,
})

const getMessageUploadLinkSucces = (payload: ApiTypes.UploadLink | null) => ({
  type: Types.GET_MESSAGE_UPLOAD_LINK_SUCCESS,
  payload
})

const setAttachmentRequest = (payload: ApiTypes.Messages.Attachment) => ({
  type: Types.SET_MESSAGE_ATTACHMENT_REQUEST,
  payload,
})

const setAttachmentSuccess = () => ({
  type: Types.SET_MESSAGE_ATTACHMENT_SUCCESS,
})

const linkMessageRequest = (payload: ApiTypes.Messages.Like) => ({
  type: Types.LIKE_MESSAGE_REQUEST,
  payload,
})

const linkMessageSuccess = () => ({
  type: Types.LIKE_MESSAGE_SUCCESS,
})

const hideMessageRequest = (payload: ApiTypes.Messages.Hide) => ({
  type: Types.HIDE_MESSAGE_REQUEST,
  payload,
})

const hideMessageSuccess = () => ({
  type: Types.HIDE_MESSAGE_SUCCESS,
})

const hideCommentRequest = (payload: ApiTypes.Messages.Hide) => ({
  type: Types.HIDE_COMMENT_REQUEST,
  payload,
})

const hideCommentSuccess = () => ({
  type: Types.HIDE_COMMENT_SUCCESS,
})

const linkCommnetRequest = (payload: ApiTypes.Messages.Like) => ({
  type: Types.LIKE_COMMENT_REQUEST,
  payload,
})

const linkCommentSuccess = () => ({
  type: Types.LIKE_COMMENT_SUCCESS,
})

const getLikesForMessageRequest = (payload: ApiTypes.Messages.Like) => ({
  type: Types.GET_LIKES_FOR_MESSAGE_REQUEST,
  payload,
})

const getLikesForMessageSuccess = (payload: ApiTypes.Messages.LikesInfoData) => ({
  type: Types.GET_LIKES_FOR_MESSAGE_SUCCESS,
  payload,
})

const getLikesForCommentRequest = (payload: ApiTypes.Messages.Like) => ({
  type: Types.GET_LIKES_FOR_COMMENT_REQUEST,
  payload,
})

const getLikesForCommentSuccess = (payload: ApiTypes.Messages.LikesInfoData) => ({
  type: Types.GET_LIKES_FOR_COMMENT_SUCCESS,
  payload,
})

const getMoreMessagesRequest = () => ({
  type: Types.GET_MORE_MESSAGES_REQUEST,
})

const getMoreMessagesSucces = (payload: CommonTypes.HubTypes.CurrentHub[]) => ({
  type: Types.GET_MORE_MESSAGES_SUCCESS,
  payload
})

const getMoreMessagesFailed = () => ({
  type: Types.GET_MORE_MESSAGES_FAILED,
})

const getMoreMessagesFromHubRequest = (payload: ApiTypes.Messages.MessagesFromHub) => ({
  type: Types.GET_MORE_MESSAGES_FROM_HUB_REQUEST,
  payload,
})

const getMoreMessagesFromHubSuccess = (payload: {
  hub: string
  messages: ApiTypes.Messages.Message[]
}) => ({
  type: Types.GET_MORE_MESSAGES_FROM_HUB_SUCCESS,
  payload
})

const getMoreMessagesFromHubFailed = () => ({
  type: Types.GET_MORE_MESSAGES_FROM_HUB_FAILED,
})

const getMessagesByIdFromHubRequest = (payload: ApiTypes.Messages.MessagesById) => ({
  type: Types.GET_MESSAGE_BY_ID_FROM_HUB_REQUEST,
  payload,
})

const getMessagesByIdFromHubSuccess = (payload: {
  hub: string
  messages: ApiTypes.Messages.Message
}) => ({
  type: Types.GET_MESSAGE_BY_ID_FROM_HUB_SUCCESS,
  payload
})

const getMessagesByIdFromHubFailed = () => ({
  type: Types.GET_MESSAGE_BY_ID_FROM_HUB_FAILED,
})

const resetMessageById = () => ({
  type: Types.RESET_MESSAGE_BY_ID,
})

const cleanAllMessages = () => ({
  type: Types.CLEAN_ALL_MESSAGES,
})

export default {
  getMessagesRequest,
  getMessagesSuccess,
  getCurrentHubRequest,
  getCurrentHubSuccess,
  getCurrentHubFailed,
  postMessageRequest,
  postMessageSucces,
  getMessagesFromHubRequest,
  getMessagesFromHubSuccess,
  getMessagesFromHubFailed,
  deleteMessageRequest,
  deleteMessageSucces,
  editMessageRequest,
  editMessageSucces,
  postCommentRequest,
  postCommentSucces,
  editCommentRequest,
  editCommentSucces,
  deleteCommentRequest,
  deleteCommentSucces,
  getMessageUploadLinkRequest,
  getMessageUploadLinkSucces,
  setAttachmentRequest,
  setAttachmentSuccess,
  linkMessageRequest,
  linkMessageSuccess,
  linkCommnetRequest,
  linkCommentSuccess,
  getLikesForMessageRequest,
  getLikesForMessageSuccess,
  getLikesForCommentRequest,
  getLikesForCommentSuccess,
  getMoreMessagesRequest,
  getMoreMessagesSucces,
  getMoreMessagesFailed,
  getMoreMessagesFromHubRequest,
  getMoreMessagesFromHubSuccess,
  getMoreMessagesFromHubFailed,
  getMessagesByIdFromHubRequest,
  getMessagesByIdFromHubSuccess,
  resetMessageById,
  getMessagesByIdFromHubFailed,
  hideMessageRequest,
  hideMessageSuccess,
  hideCommentRequest,
  hideCommentSuccess,
  reportMessageHubRequest,
  reportMessageCentralRequest,
  reportMessageSucces,
  cleanAllMessages,
}