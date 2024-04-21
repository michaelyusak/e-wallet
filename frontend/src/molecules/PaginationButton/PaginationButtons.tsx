import React from 'react';
import paginationLeftIcon from '../../assets/img/pagination-left.png';
import paginationRightIcon from '../../assets/img/pagination-right.png';
import * as S from './PaginationButtons.Styles';

type PaginationButtonsProps = {
  setPage: (diff: number) => void;
};

const PaginationButtons = ({
  setPage,
}: PaginationButtonsProps): React.ReactElement => {
  return (
    <S.Buttons>
      <button onClick={() => setPage(-1)}>
        <img src={paginationLeftIcon} alt=""></img>
      </button>

      <button onClick={() => setPage(1)}>
        <img src={paginationRightIcon} alt=""></img>
      </button>
    </S.Buttons>
  );
};

export default PaginationButtons;
