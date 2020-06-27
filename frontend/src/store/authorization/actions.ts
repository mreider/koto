
import { ApiTypes } from './../../types'

export enum Types {
  LOGIN_REQUESTED = 'LOGIN_REQUESTED',
  LOGIN_SUCCESS = 'LOGIN_SUCCESS',
  LOGIN_FAILED = 'LOGIN_FAILED',
  RESET_LOGIN_FAILED_MESSAGE = 'RESET_LOGIN_FAILED_MESSAGE',
}

export const loginRequested = (payload: ApiTypes.Login) => ({
  type: Types.LOGIN_REQUESTED,
  payload
})

export const loginSucces = () => ({
  type: Types.LOGIN_SUCCESS
})

export const loginFailed = (payload: string) => ({
  type: Types.LOGIN_FAILED,
  payload
})

export const resetLoginFailedMessage = () => ({
  type: Types.RESET_LOGIN_FAILED_MESSAGE
})

export default {
  loginRequested,
  loginSucces,
  loginFailed,
  resetLoginFailedMessage,
}