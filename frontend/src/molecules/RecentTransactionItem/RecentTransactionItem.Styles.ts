import styled from 'styled-components';

export const IncomeCard = styled.div`
  display: flex;
  padding: 10px 10px 10px 24px;
  justify-content: space-between;
  align-items: center;
  height: 103px;
  border: solid #ede6e7 1px;
  border-radius: 15px;

  img {
    width: 39px;
  }

  p:nth-child(2) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 40%;
  }

  p:nth-child(3) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  p:nth-child(4) {
    color: #33a720;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  @media (max-width: 1050px) {
    padding: 8px;

    img {
      width: 28px;
    }

    p:nth-child(2),
    p:nth-child(3),
    p:nth-child(4) {
      line-height: 24px;
      font-size: 16px;
    }

    p:nth-child(2) {
      width: 30%;
    }

    p:nth-child(3) {
      width: 27%;
    }
  }
`;

export const ExpenseCard = styled.div`
  display: flex;
  padding: 10px 10px 10px 24px;
  justify-content: space-between;
  align-items: center;
  height: 103px;
  border: solid #ede6e7 1px;
  border-radius: 15px;

  img {
    width: 39px;
  }

  p:nth-child(2) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 40%;
  }

  p:nth-child(3) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  p:nth-child(4) {
    color: #f60707;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  @media (max-width: 1050px) {
    padding: 8px;

    img {
      width: 28px;
    }

    p:nth-child(2),
    p:nth-child(3),
    p:nth-child(4) {
      line-height: 24px;
      font-size: 16px;
    }

    p:nth-child(2) {
      width: 30%;
    }

    p:nth-child(3) {
      width: 27%;
    }
  }
`;

export const RecentTransactionCard = styled.div<{$type:'income' | 'expense'}>`
  display: flex;
  padding: 10px 10px 10px 24px;
  justify-content: space-between;
  align-items: center;
  height: 103px;
  border: solid 1px ${(props) => props.$type === 'expense' ? '#ede6e7' : props.$type === 'income' && '#ede6e7'};
  border-radius: 15px;

  img {
    width: 39px;
  }

  p:nth-child(2) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 40%;
  }

  p:nth-child(3) {
    color: #282828;
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  p:nth-child(4) {
    color: ${(props) => props.$type === 'expense' ? '#f60707' : props.$type === 'income' && '#33a720'};
    line-height: 27px;
    font-size: 18px;
    font-weight: 600;
    width: 20%;
  }

  @media (max-width: 1050px) {
    padding: 8px;

    img {
      width: 28px;
    }

    p:nth-child(2),
    p:nth-child(3),
    p:nth-child(4) {
      line-height: 24px;
      font-size: 16px;
    }

    p:nth-child(2) {
      width: 30%;
    }

    p:nth-child(3) {
      width: 27%;
    }
  }
`;