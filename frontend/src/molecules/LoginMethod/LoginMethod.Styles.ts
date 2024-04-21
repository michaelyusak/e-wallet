import styled from 'styled-components';

export const Methods = styled.div`
  margin-top: 13%;

  p {
    text-align: center;

    line-height: 24px;
    font-size: 16px;
    font-weight: 500;
    color: #b5b5b5;
  }

  div {
    display: flex;
    justify-content: center;
    gap: 6%;
    margin-top: 13%;

    button {
      height: 44px;
      border: none;
      outline: none;
      background-color: transparent;
      cursor: pointer;

      img {
        width: 41.46px;
      }
    }
  }

  @media (max-width: 710px) {
    div {
      margin-top: 17px;
      gap: 3%;
    }
  }
`;
