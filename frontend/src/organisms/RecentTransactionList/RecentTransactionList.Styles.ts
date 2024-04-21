import styled from 'styled-components';

export const RecentTransactionList = styled.div`
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: 329px;
`;

export const NoTransactionStatement = styled.div`
  margin: 0 auto;
  margin-top: 267px;

  h3 {
    text-align: center;
    color: #282828;
    line-height: 33px;
    font-size: 22px;
    font-weight: 600;
  }

  p {
    margin-top: 1px;
    text-align: center;
    color: #95999e;
    line-height: 18px;
    font-size: 12px;
    font-weight: 600;
  }
`;
