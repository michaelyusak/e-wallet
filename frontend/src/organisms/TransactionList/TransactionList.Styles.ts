import styled from 'styled-components';

export const TransactionListTable = styled.table`
  width: 100%;
  border-collapse: collapse;

  th {
    padding: 12px 24px;
    width: 25%;
  }

  th:nth-child(1),
  th:nth-child(2),
  th:nth-child(3) {
    text-align: left;
  }

  th:nth-child(4) {
    text-align: right;
  }

  thead {
    th {
      line-height: 16px;
      font-size: 12px;
      font-weight: 700;
      color: #4a5568;
    }
  }
`;
