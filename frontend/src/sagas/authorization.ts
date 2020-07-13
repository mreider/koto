import { put } from 'redux-saga/effects'
import Actions from '@store/actions'
import { ApiTypes } from '../types/index'
import { API } from '@services/api' 

export function* watchlogin(action: { type: string, payload: ApiTypes.Login }) {
  const response = yield API.authorization.login(action.payload)

  if (response.status === 200) {
    localStorage.setItem('kotoIsLogged', 'true')
    yield put(Actions.authorization.loginSucces())
    yield put(Actions.authorization.getAuthTokenRequest())
    yield put(Actions.profile.getProfileRequest())
  } else {
    yield put(Actions.authorization.loginFailed(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchlogout() {
  const response = yield API.authorization.logout()
  if (response.status === 200) {
    yield put(Actions.authorization.logoutSucces())
  }
}

export function* watchGetAuthToken() {
  const response = yield API.authorization.getAuthToken()

  if (response.status === 200) {
    localStorage.setItem('kotoAuthToken', JSON.stringify(response.data?.token))
    yield put(Actions.authorization.getAuthTokenSucces(response.data?.token))
  } else {
    yield put(Actions.authorization.loginFailed(response?.error?.response?.data?.msg || 'Server error'))
  }
}