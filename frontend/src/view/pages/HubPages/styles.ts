import styled from 'styled-components'
import Container from '@material-ui/core/Container'
import FormControl from '@material-ui/core/FormControl'
import Button from '@material-ui/core/Button'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import { Link } from 'react-router-dom'

export const HubsListContainer = styled.main`
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  position: relative;
  max-width: 1170px;
  width: 100%;
  margin: 65px auto 0;
  padding: 30px 15px 30px;

  @media (max-width: 770px) {
    flex-wrap: wrap;
    margin-top: 50px;
    padding: 10px 15px 15px;
}
`

export const ContainerStyled = styled(Container)`
  && {
    display: flex;
    flex-wrap: wrap;
    align-content: center;
    justify-content: center;
    margin-top: 60px;
    padding-bottom: 30px;
  }
`

export const HubListWrapper = styled.div`
  width: 100%;

  @media (max-width: 600px) {
    padding: 0 15px;
  }
`

export const FormWrapper = styled.form`
  width: 400px;
  margin: 0 auto;
`

export const FormControlStyled = styled(FormControl)`
  && {
    margin: 0 0 15px; 
    width: 100%;
  }
`

export const ButtonStyled = styled(Button)`
  && {
    height: 42px;
    min-width: 200px;
  }

  &.yellow {
    && {
      background: #f2c641;
    }
  }  
`

export const TitleWrapper = styled.div`
  text-align: center;
  margin-bottom: 20px;
`

export const FormControlLabelStyled = styled(FormControlLabel)`
  && * {
    font-size: 15px;
  }

  && {
    margin-left: 0px;
  }
`

export const TableWrapper = styled.div`
  margin-top: 20px;
  width: 100%;
`

export const TableCellStyled = styled(TableCell)`
  padding: 8px 16px;
`

export const TableHeadStyled = styled(TableHead)`
  && * {
    font-weight: bold;
  }
`

export const ApproveButton = styled(Button)`
  && {
    background: #34c242;

    &:hover {
      background: #32ab3d;
    }
  }
`

export const PendingTextWrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
`

export const PendingText = styled.span`
  margin-left: 5px;
`

export const LinkWrapper = styled.div`
  max-width: 400px;
  width: 100%;
  text-align: right;
  margin-top: 30px;
`

export const LinkStyled = styled(Link)`
  color: #1976d2;
  font-size: 1rem;
  text-decoration: none;
  display: block;

  &:hover {
    text-decoration: underline;
  }
`

export const ALinkStyled = styled.a`
  color: #1976d2;
  font-size: 1rem;
  text-decoration: none;
  display: block;

  &:hover {
    text-decoration: underline;
  }
`