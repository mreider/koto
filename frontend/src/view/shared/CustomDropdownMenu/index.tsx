import React, { useState } from 'react'
import { connect } from 'react-redux'
import { history } from '@view/routes'
import Actions from '@store/actions'
import Avatar from '@material-ui/core/Avatar'
import { getAvatarUrl } from '@services/avatarUrl'
import selectors from '@selectors/index'
import { StoreTypes } from 'src/types'
import ListItemText from '@material-ui/core/ListItemText'
import ExitToAppIcon from '@material-ui/icons/ExitToApp'
import ClickAwayListener from '@material-ui/core/ClickAwayListener'
import PersonIcon from '@material-ui/icons/Person'
import {
  AvatarWrapper,
  ListItemIconStyled,
} from '../styles'

import {
  DropdownMenuWrapper,
  Dropdown,
  CustomMenuItem,
} from './styles'

interface Props {
  userId: string

  onLogout: () => void
}

const DropdownMenu2: React.FC<Props> = (props) => {
  const [isMenuOpen, openMenu] = useState(false)
  const { userId } = props

  const goToPage = (path: string) => {
    openMenu(false)
    history.push(path)
  }

  const onLogoutClick = () => {
    localStorage.clear()
    sessionStorage.clear()
    history.push('/login')
    props.onLogout()
  }

  return (
    <ClickAwayListener onClickAway={() => { openMenu(false) }}>
      <DropdownMenuWrapper>
        <AvatarWrapper onClick={() => openMenu(!isMenuOpen)}>
          <Avatar src={getAvatarUrl(userId)} />
        </AvatarWrapper>
        {isMenuOpen && <Dropdown>
          <CustomMenuItem onClick={() => goToPage('/settings')}>
            <ListItemIconStyled>
              <PersonIcon fontSize="small" />
            </ListItemIconStyled>
            <ListItemText primary="View profile" />
          </CustomMenuItem>
          <CustomMenuItem className="logout" onClick={onLogoutClick}>
            <ListItemIconStyled>
              <ExitToAppIcon fontSize="small" />
            </ListItemIconStyled>
            <ListItemText primary="Log out" />
          </CustomMenuItem>
        </Dropdown>}
      </DropdownMenuWrapper>
    </ClickAwayListener>
  )
}

type StateProps = Pick<Props, 'userId'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  userId: selectors.profile.userId(state),
})

type DispatchProps = Pick<Props, 'onLogout'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onLogout: () => dispatch(Actions.authorization.logoutRequest()),
})

export default connect(mapStateToProps, mapDispatchToProps)(DropdownMenu2)
