import styled from 'styled-components';

export const WalletBalance = styled.div`
  background-color: #f9f9f9;
  padding: 26px 22px 26px 22px;
  border-radius: 15px;
  width: 40%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  h3 {
    color: #4d47c3;
    line-height: 33px;
    font-size: 22px;
    font-weight: 600;
  }

  button {
    background-color: transparent;
    border: none;
    outline: none;
    text-align: left;
    width: fit-content;
    cursor: pointer;
  }

  :nth-child(2) {
    line-height: 42px;
    font-size: 28px;
    font-weight: 600;
  }

  :nth-child(3) {
    color: #4d47c3;
    line-height: 21px;
    font-size: 14px;
    font-weight: 600;
  }

  @media (max-width: 1050px) {
    width: 96%;
    height: 187px;
  }
`;

export const IncomeFlowCard = styled.div`
  background-color: #eafcef;
  width: 35%;
  border-radius: 15px;
  padding: 26px 23px 31px 19px;
  border: solid #c9ffd8 1px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  img:nth-child(1) {
    width: 28px;
  }

  div:nth-child(2) {
    display: flex;
    justify-content: space-between;
    align-items: center;

    p {
      line-height: 42px;
      font-size: 28px;
      font-weight: 600;
      margin: auto 0;
    }

    img {
      width: 66px;
      height: 47px;
    }
  }

  p:nth-child(3) {
    line-height: 24px;
    font-size: 16px;
    font-weight: 600;
    color: #33a720;
  }

  @media (max-width: 1050px) {
    width: 96%;
    height: 187px;
  }
`;

export const ExpenseFlowCard = styled.div`
  background-color: #feefe7;
  width: 35%;
  border-radius: 15px;
  padding: 26px 23px 31px 19px;
  border: solid #ffddca 1px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  img:nth-child(1) {
    width: 28px;
  }

  div:nth-child(2) {
    display: flex;
    justify-content: space-between;
    align-items: center;

    p {
      line-height: 42px;
      font-size: 28px;
      font-weight: 600;
      margin: auto 0;
    }

    img {
      width: 66px;
      height: 47px;
    }
  }

  p:nth-child(3) {
    line-height: 24px;
    font-size: 16px;
    font-weight: 600;
    color: #f60707;
  }

  @media (max-width: 1050px) {
    width: 96%;
    height: 187px;
  }
`;
