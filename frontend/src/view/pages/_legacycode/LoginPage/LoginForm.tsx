import React, { FormEvent, useState, useEffect } from 'react'
import Typography from '@material-ui/core/Typography'
import Visibility from '@material-ui/icons/Visibility'
import VisibilityOff from '@material-ui/icons/VisibilityOff'
import InputLabel from '@material-ui/core/InputLabel'
import OutlinedInput from '@material-ui/core/OutlinedInput'
import InputAdornment from '@material-ui/core/InputAdornment'
import IconButton from '@material-ui/core/IconButton'
import CircularProgress from '@material-ui/core/CircularProgress'
import FormHelperText from '@material-ui/core/FormHelperText'
import { validate } from '@services/validation'
import { FooterMenu } from '@view/shared/FooterMenu'
import logo from './../../../assets/images/logo-2.png'
import { useLastLocation } from 'react-router-last-location'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Checkbox from '@material-ui/core/Checkbox'
import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank'
import CheckBoxIcon from '@material-ui/icons/CheckBox'
import {
  FormWrapper,
  FormControlStyled,
  ButtonStyled,
  ContainerStyled,
  Header,
  Logo,
  CheckBoxWrapper
} from './styles'
import { ApiTypes } from 'src/types'
import { RouteComponentProps } from 'react-router-dom'

type FieldsType = 'username' | 'password' | ''

export interface Props extends RouteComponentProps {
  loginErrorMessage: string
  isLogged: boolean

  onLogin: (data: ApiTypes.Login) => void
  resetLoginFailedMessage: () => void
}

export const LoginForm = (props) => {
  const [username, onEmailChange] = useState<string>('')
  const [password, onPasswordChange] = useState<string>('')
  const [isRememberedMe, onRememberMeChange] = useState<boolean>(false)
  const [isPasswordVisible, onPasswordOpen] = useState<boolean>(false)
  const [errorMessage, setErrorMessage] = useState<string>('')
  const [noValideField, setNoValideField] = useState<FieldsType>('')
  const [isRequest, setRequest] = useState<boolean>(false)
  const { loginErrorMessage, isLogged, onLogin, location, history } = props

  const onValidate = (): boolean => {
    if (!validate.isUserNameValid(username)) {
      setErrorMessage('Your username contains invalid characters, allowed: a-z A-Z 0-9 @ . -')
      setNoValideField('username')
      return false
    }

    if (!validate.isPasswordValid(password)) {
      setErrorMessage('Password is incorrect')
      setNoValideField('password')
      return false
    }

    return true
  }

  const onFormSubmit = (event: FormEvent) => {
    props.resetLoginFailedMessage()
    event.preventDefault()
    if (!onValidate()) return

    setRequest(true)
    setErrorMessage('')
    setNoValideField('')

    onLogin({ 
      name: username, 
      password,
      remember_me: isRememberedMe,
    })
  }

  const checkExcludedRoutes = (path?: string) => {
    const excludedRoutes = ['reset-password', 'forgot-password', 'docs', 'no-hubs', 'registration']
    if (!path) return false
    return excludedRoutes.some(item => path.indexOf(item) !== -1)
  }

  const lastLocation = useLastLocation()
  const lastLoactionPathname = lastLocation?.pathname

  useEffect(() => {
    if (loginErrorMessage !== '') {
      setErrorMessage(loginErrorMessage)
      setRequest(false)
    }

    if (isLogged === true) {
      if (location.pathname !== lastLoactionPathname) {
        if (lastLoactionPathname === '/registration' || lastLoactionPathname === '/resend-confirm-email') {
          history.push('/friends/invitations')
        }
        if (checkExcludedRoutes(lastLoactionPathname)) {
          history.push('/messages')
        } else {
          history.push(lastLoactionPathname ? `${lastLoactionPathname}` : '/messages')
        }
      }
      setRequest(false)
    }
  }, [loginErrorMessage, isLogged, onLogin, location, history, lastLoactionPathname])

  return (
    <ContainerStyled maxWidth="sm">
      <Header>
        <Logo src={logo} />
        <Typography variant="subtitle1" gutterBottom>Peacenik is an ad-free, friendly, distributed social network.</Typography>
      </Header>
      <FormWrapper onSubmit={onFormSubmit}>
        <FormControlStyled variant="outlined">
          <InputLabel
            htmlFor="username"
            color={(noValideField === 'username') ? 'secondary' : 'primary'}
          >User Name</InputLabel>
          <OutlinedInput
            id="username"
            type={'text'}
            value={username}
            error={(noValideField === 'username') ? true : false}
            onChange={(event) => onEmailChange(event.currentTarget.value.trim())}
            labelWidth={85}
          />
        </FormControlStyled>
        <FormControlStyled variant="outlined">
          <InputLabel
            htmlFor="password"
            color={(noValideField === 'password') ? 'secondary' : 'primary'}
          >Password</InputLabel>
          <OutlinedInput
            id="password"
            type={isPasswordVisible ? 'text' : 'password'}
            value={password}
            onChange={(event) => onPasswordChange(event.currentTarget.value.trim())}
            error={(noValideField === 'password') ? true : false}
            labelWidth={70}
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={() => onPasswordOpen(!isPasswordVisible)}
                  onMouseDown={() => onPasswordOpen(!isPasswordVisible)}
                  edge="end"
                >
                  {isPasswordVisible ? <Visibility /> : <VisibilityOff />}
                </IconButton>
              </InputAdornment>
            } />
        </FormControlStyled>
        <CheckBoxWrapper>
          <FormControlLabel
            control={
              <Checkbox
                checked={isRememberedMe}
                onChange={(event) => onRememberMeChange(event.target.checked)}
                icon={<CheckBoxOutlineBlankIcon fontSize="small" />}
                checkedIcon={<CheckBoxIcon fontSize="small" />}
                name="rememberMe"
                color="primary"
              />
            }
            label="Remember me"
          />
        </CheckBoxWrapper>
        <ButtonStyled
          variant="contained"
          size="large"
          color="primary"
          type="submit"
          onClick={onFormSubmit}
        >
          {isRequest ? <CircularProgress size={25} color={'inherit'} /> : 'Login'}
        </ButtonStyled>
        {errorMessage && <FormHelperText error>{errorMessage}</FormHelperText>}
      </FormWrapper>
      <FooterMenu menuItems={[
        {
          title: 'Forgotten password',
          to: '/forgot-password'
        },
        {
          title: 'Forgotten username',
          to: '/forgot-username'
        },
        {
          title: 'Register for Peacenik',
          to: '/registration',
        },
        {
          title: 'About Peacenik',
          href: 'https://docs.koto.at',
        },
        {
          title: 'End User License Agreement (EULA)',
          to: '/docs/code-of-conduct'
        },
        {
          title: 'Contact us',
          href: 'https://docs.koto.at/#/help',
        },
      ]} />
    </ContainerStyled>
  )
} 
