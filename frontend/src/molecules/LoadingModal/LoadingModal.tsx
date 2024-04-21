import React from 'react';
import loadingGif from '../../assets/img/loading-circle.gif';
import * as S from './LoadingModal.Styles';

const LoadingModal = (): React.ReactElement => {
  return (
    <S.ModalContainer>
      <S.ModalUnderlay></S.ModalUnderlay>

      <S.ModalCard>
        <img src={loadingGif} alt=""></img>
      </S.ModalCard>
    </S.ModalContainer>
  );
};

export default LoadingModal;
