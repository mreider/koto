import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import Friends from './Friends'
import Invitations from './Invitations'
import FriendInvite from './FriendInvite'
import FriendTabs from './Tabs'
import { Switch, Route, RouteComponentProps } from 'react-router-dom'
import { PageLayout } from '@view/shared/PageLayout'
import { PageNavigation } from './PageNavigation'
import { FriendsWrapper } from './styles'

interface Props extends RouteComponentProps {
  onGetFriends: () => void
  onGetInvitations: () => void
}

const FriendsPage: React.FC<Props> = (props) => {
  const { onGetFriends, onGetInvitations } = props

  useEffect(() => {
    onGetFriends()
    onGetInvitations()
  })

  return (
    <PageLayout>
      <PageNavigation />
      <FriendsWrapper>
        <FriendTabs />
        <Switch>
          <Route path="/friends" exact component={Friends} />
          <Route path="/friends/all" exact component={Friends} />
          <Route path="/friends/invitations" exact component={Invitations} />
          <Route path="/friends/invite" exact component={FriendInvite} />
        </Switch>
      </FriendsWrapper>
    </PageLayout>
  )
}

type DispatchProps = Pick<Props, 'onGetFriends' | 'onGetInvitations'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onGetFriends: () => dispatch(Actions.friends.getFriendsRequest()),
  onGetInvitations: () => dispatch(Actions.friends.getInvitationsRequest()),
})

export default connect(null, mapDispatchToProps)(FriendsPage)