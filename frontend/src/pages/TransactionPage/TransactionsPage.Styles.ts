import styled from 'styled-components';

export const FullPage = styled.section`
  display: flex;

  @media (max-width: 1050px) {
    height: fit-content;
  }
`;

export const TransactionsContent = styled.div`
  width: 100%;
  padding: 20px 60px 26px 65px;
  margin-left: 280px;
`;

export const Hero = styled.div`
  margin-top: 24px;

  button {
    border: none;
    outline: none;
    background-color: transparent;
    cursor: pointer;
  }

  h1 {
    line-height: 63px;
    font-size: 42px;
    font-weight: 600;
    color: #282828;
  }

  p {
    line-height: 24px;
    font-size: 16px;
    font-weight: 600;
    color: #95999e;
  }
`;

export const Transactions = styled.div`
  margin-top: 32px;
`;

export const Selectors = styled.div`
  display: flex;
  justify-content: right;
  gap: 20px;

  > div {
    align-items: center;

    p {
      line-height: 16px;
      font-size: 12px;
      font-weight: 700;
    }

    div:nth-child(2) {
      button {
        line-height: 16px;
        font-size: 12px;
        font-weight: 700;
        cursor: pointer;
      }

      > button {
        background-color: #ede6e7;
        border: none;
        padding: 7px 5px;
        border-radius: 10px;
      }

      > div {
        background-color: #ede6e7;
        padding: 10px 16px;
        border-radius: 10px;
        gap: 10px;

        button {
          background-color: transparent;
          border: none;
          outline: none;
        }
      }
    }
  }

  > div:nth-child(1) > div:nth-child(2) {
    > button {
      width: 89px;
    }

    > div {
      width: 89px;
    }
  }

  > div:nth-child(2) > div:nth-child(2) {
    > button {
      width: 160px;
    }

    > div {
      width: 160px;
    }
  }
`;

export const TransactionListContainer = styled.div`
  border: solid #e2e8f0 1px;
  margin-top: 10px;
  padding: 5px 12px;
  border-radius: 10px;
  height: 470px;
`;

export const AddTransactionButtonsContainer = styled.div`
  display: flex;
  margin-top: 10px;
  justify-content: flex-end;
  gap: 24px;
`;

export const Status = styled.p`
  margin-left: 280px;
`;
