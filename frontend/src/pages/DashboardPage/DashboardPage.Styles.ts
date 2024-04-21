import styled from 'styled-components';

export const FullPage = styled.section`
  display: flex;

  @media (max-width: 1050px) {
    height: fit-content;
  }
`;

export const DashboardContent = styled.section`
  width: 100%;
  padding: 30px 60px 26px 65px;
  margin-left: 280px;

  h2:nth-child(1) {
    color: #282828;
    line-height: 42px;
    font-size: 26px;
    font-weight: 600;
  }

  @media (max-width: 710px) {
    margin-left: 74px;
    padding: 29px 22px 29px 22px;

    h2:nth-child {
      line-height: 36px;
      font-size: 24px;
    }
  }
`;

export const WalletInformation = styled.div`
  display: flex;
  gap: 23px;
  margin-top: 2%;
  height: 187px;

  @media (max-width: 1050px) {
    flex-direction: column;
    height: fit-content;
  }
`;

export const RecentTransactions = styled.div`
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-top: 7%;

  @media (max-width: 710px) {
    margin-top: 22%;
  }
`;

export const RecentTransactionsHeader = styled.div`
  h2 {
    color: #282828;
    line-height: 42px;
    font-size: 28px;
    font-weight: 600;
  }

  p {
    color: #95999e;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
  }

  @media (max-width: 710px) {
    h2 {
      line-height: 36px;
      font-size: 24px;
    }

    p {
      line-height: 24px;
      font-size: 16px;
    }
  }
`;

export const Status = styled.p`
  margin-left: 280px;
`;
