import React from 'react'
import { Link } from 'react-router-dom'
import { ApiTypes } from 'src/types'
import { getAvatarUrl } from '@services/avatarUrl'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import { ButtonOutlined, ButtonContained } from '@view/shared/styles'
import {
  MemberWrapper,
  MemberAvatar,
  MemberName,
  MemberHeaderSidebar,
  MemberNameWrapperSidebar,
  MemberMessageSidebar,
  MemberButtonsWrapperSidebar,
} from './styles'

interface Props extends ApiTypes.Groups.Invite {
  calback?: (data: ApiTypes.Groups.ConfirmDenyInvite) => void
  onConfirmInvite: (data: ApiTypes.Groups.ConfirmDenyInvite) => void
  onDenyInvite: (data: ApiTypes.Groups.ConfirmDenyInvite) => void
}

const MemberInvited: React.FC<Props> = (props) => {
  const {
    group_id,
    inviter_id,
    invited_name,
    invited_full_name,
    message,
    invited_id,
    onDenyInvite,
    onConfirmInvite,
    calback,
  } = props

  const data = {
    group_id,
    inviter_id,
    invited_id,
  }

  const onDeny = () => {
    onDenyInvite(data)
    calback && calback(data)
  }

  const onConfirm = () => {
    onConfirmInvite(data)
    calback && calback(data)
  }

  return (
    <MemberWrapper className="sidebar">
      <MemberHeaderSidebar>
        <Link to={`/profile/user?id=${invited_id}`}>
          <MemberAvatar src={getAvatarUrl(invited_id)} className="sidebar" />
        </Link>
        <MemberNameWrapperSidebar>
          <MemberName
            className="sidebar"
            to={`/profile/user?id=${invited_id}`}>
            {invited_full_name || invited_name}
          </MemberName>
          <MemberMessageSidebar>{message}</MemberMessageSidebar>
        </MemberNameWrapperSidebar>
      </MemberHeaderSidebar>
      <MemberButtonsWrapperSidebar>
        <ButtonOutlined onClick={onDeny} className="small gray">Ignore</ButtonOutlined>
        <ButtonContained onClick={onConfirm} className="small">Approve</ButtonContained>
      </MemberButtonsWrapperSidebar>
    </MemberWrapper>
  )
}

type DispatchProps = Pick<Props, 'onConfirmInvite' | 'onDenyInvite'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onConfirmInvite: (data: ApiTypes.Groups.ConfirmDenyInvite) => dispatch(Actions.groups.confirmInviteRequest(data)),
  onDenyInvite: (data: ApiTypes.Groups.ConfirmDenyInvite) => dispatch(Actions.groups.denyInviteRequest(data)),
})

export default connect(null, mapDispatchToProps)(MemberInvited)
