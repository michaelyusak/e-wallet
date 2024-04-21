import styled from 'styled-components';

export const FullPage = styled.section`
  height: 100vh;
`;

export const Header = styled.header`
  padding: 31px 42px;

  @media (max-width: 710px) {
    padding: 38px 16px;
  }
`;

export const Logo = styled.h1`
  font-weight: 600;
  font-size: 20px;
  line-height: 30px;
  color: #4d47c3;
`;

export const PageContent = styled.div`
  display: flex;
  justify-content: space-between;

  @media (max-width: 710px) {
    flex-direction: column-reverse;
  }
`;

export const PageDescription = styled.div`
  margin-top: 14%;
  margin-left: 10%;
  height: fit-content;

  p:nth-child(1) {
    line-height: 52.5px;
    font-size: 35px;
    font-weight: 500;

    max-width: 240px;

    b {
      line-height: 75px;
      font-size: 50px;
      font-weight: 600;
    }
  }

  p:nth-child(2) {
    margin-top: 13%;
    line-height: 30px;
    font-weight: 400;
    font-size: 16px;
    max-width: 300px;

    button {
      border: none;
      outline: none;
      background-color: transparent;
      padding: 0px 10px;

      color: #4d47c3;
      line-height: 24px;
      font-size: 16px;
      font-weight: 600;

      cursor: pointer;
    }
  }

  @media (max-width: 710px) {
    display: none;
  }
`;

export const LoginImg = styled.img`
  width: 20%;
  height: 600px;
  margin-top: 13%;
  margin-bottom: 0.5%;
  object-fit: cover;

  @media (max-width: 710px) {
    display: none;
  }
`;

export const SignInForm = styled.div`
  display: flex;
  flex-direction: column;
  margin-right: 10%;
  margin-top: 6%;
  width: 26%;
  height: fit-content;

  h2 {
    line-height: 45px;
    font-size: 30px;
    font-weight: 500;
    margin-bottom: 7%;
  }

  @media (max-width: 710px) {
    margin: 0 auto;
    margin-top: 58px;
    width: 91%;
  }
`;

export const MobileDescription = styled.div`
  display: none;

  @media (max-width: 710px) {
    display: flex;
    margin-top: 30px;
    margin-bottom: 125px;

    p {
      line-height: 30px;
      font-weight: 400;
      font-size: 16px;

      max-width: 300px;

      button {
        border: none;
        outline: none;
        background-color: transparent;
        padding: 0px 10px;

        color: #4d47c3;
        line-height: 24px;
        font-size: 16px;
        font-weight: 600;

        cursor: pointer;
      }
    }
  }
`;
