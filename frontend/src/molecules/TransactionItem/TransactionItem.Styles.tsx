import styled from 'styled-components';

export const TransactionItemRow = styled.tr<{
  $transactionType: 'income' | 'expense';
}>`
  th:nth-child(4) {
    color: ${(props) =>
      props.$transactionType === 'income' ? '#33A720' : '#F60707'};
  }

  th {
    line-height: 20px;
    font-size: 14px;
    font-weight: 500;
    color: #2d3748;
    border-top: #e2e8f0 solid 1px;
    height: 52px;
  }
`;
