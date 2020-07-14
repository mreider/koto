import { Types } from './actions'
import { NodeTypes, ApiTypes } from '../../types'

export interface State {
  messageTokens: NodeTypes.CurrentNode[]
  currentNode: NodeTypes.CurrentNode
  isMessagePostedSuccess: boolean
  messages: ApiTypes.Messages.Message[]
}

const initialState: State = {
  messageTokens: [],
  currentNode: {
    host: '',
    token: '',
  },
  isMessagePostedSuccess: false,
  messages: []
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case Types.GET_MESSAGES_SUCCESS: {
      return {
        ...state, ...{ messageTokens: action.payload }
      }
    }
    case Types.GET_CURRENT_NODE_SUCCESS: {
      return {
        ...state, ...{ currentNode: action.payload }
      }
    }
    case Types.POST_MESSAGE_SUCCESS: {
      return {
        ...state, ...{ isMessagePostedSuccess: action.payload }
      }
    }
    case Types.GET_MESSAGES_FROM_NODE_SUCCESS: {
      return {
        ...state, ...{ 
          messages: [...state.messages, ...action.payload]
        }
      }
    }
    default: return state
  }
}

export default reducer