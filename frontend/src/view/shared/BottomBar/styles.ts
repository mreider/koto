
import styled from 'styled-components'
import { Link } from 'react-router-dom'

export const BottomBarWrapper = styled.footer`
  background: #FFFFFF;
  box-shadow: 0px -1px 10px rgba(212, 212, 212, 0.8);
  /* position: absolute; */
  /* z-index: 1000; */
  /* bottom: 0; */
  /* left: 0; */
  width: 100%;
  padding: 20px 0px;
  margin-top: auto;

  @media (max-width: 600px){
    padding: 15px 0;
  }
`

export const CopyrightWrapper = styled.div`
  text-align: center;
  color: #A1AEC8;
  padding-top: 20px;
  border-top: 1px solid #A1AEC8;
  font-family: 'SFUITextRegular';

  @media (max-width: 600px) {
    font-size: 14px;
    padding: 15px 15px 0px;
  }
`

export const FooterWrapper = styled.div`
  padding-bottom: 20px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  @media (max-width: 600px){
    padding-bottom: 15px;
  }
  
  @media (max-width: 370px){
    padding-bottom: 5px;
  }
`

export const LogoIcon = styled.img`
  width: 100%;
  max-width: 44px;
  
  @media (max-width: 600px){
    max-width: 40px;
    max-height: 40px;
  }

  @media (max-width: 370px){
    margin-bottom: 10px;
  }
`

export const FooterMenu = styled.nav`
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
`

export const FooterMenuItem = styled.span`
  font-size: 14px;
  text-transform: uppercase;
  white-space: nowrap;

  &:after {
    display: inline-block;
    margin: 0 30px;
    width: 1px;
    height: 16px;
    position: relative;
    top: 2px;
    background: linear-gradient(0deg, #ABB7CD, #ABB7CD), #FFFFFF;  
    content: '';
  }

  &:last-child {
    &:after {
      display: none;
    }
  }


  @media (max-width: 600px) {
    font-size: 10px;

    &:after {
      margin: 0 13px;
      height: 12px;
      top: 3px;
    }
  }

  @media (max-width: 370px) {
    margin: 0px 10px 10px 0;
    &:after {
      display: none;
    }
  }
`

export const FooterMenuLink = styled(Link)`
  text-decoration: none;
  color: #262626;
  font-weight: bold;
  transition: 0.2s;
  font-family: 'SFUITextBold';

  &:hover {
    color: #A1AEC8
  }
`

export const FooterMenuALink = styled.a`
  text-decoration: none;
  color: #262626;
  font-weight: bold;
  transition: 0.2s;
  font-family: 'SFUITextBold';

  &:hover {
    color: #A1AEC8
  }
`

export const LogoWrapper = styled(Link)`
  display: flex;
`