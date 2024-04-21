import React, { useState } from 'react';

import { IWallet } from '../../interfaces/Wallet';
import * as S from './WalletInformation.Styles';
import incomeIcon from '../../assets/img/income.png';
import expenseIcon from '../../assets/img/expense.png';
import incomeGraphIcon from '../../assets/img/income-graph.png';
import expenseGraphIcon from '../../assets/img/expense-graph.png';

type walletInformationProps = {
  wallet: IWallet;
};

const WalletInformation = ({
  wallet,
}: walletInformationProps): React.ReactElement => {
  const formatter = new Intl.NumberFormat('en-US');

  const [isShowBalance, setShowBalance] = useState<boolean>(false);

  function handleSetShowBalance() {
    setShowBalance(!isShowBalance);
  }

  return (
    <>
      <S.WalletBalance>
        <h3>Balance</h3>

        <button onClick={() => handleSetShowBalance()}>
          <p>
            IDR{' '}
            {isShowBalance
              ? formatter.format(+wallet?.balance)
              : '*'.repeat(wallet?.balance.length || 1)}
          </p>
        </button>

        <p>{wallet?.number || '1000000000000'}</p>
      </S.WalletBalance>

      <S.IncomeFlowCard>
        <img src={incomeIcon} alt=""></img>

        <div>
          <p>IDR {formatter.format(+wallet?.income)}</p>

          <img src={incomeGraphIcon} alt=""></img>
        </div>

        <p>Income</p>
      </S.IncomeFlowCard>

      <S.ExpenseFlowCard>
        <img src={expenseIcon} alt=""></img>

        <div>
          <p>IDR {formatter.format(+wallet?.expense)}</p>

          <img src={expenseGraphIcon} alt=""></img>
        </div>

        <p>Expense</p>
      </S.ExpenseFlowCard>
    </>
  );
};

export default WalletInformation;
