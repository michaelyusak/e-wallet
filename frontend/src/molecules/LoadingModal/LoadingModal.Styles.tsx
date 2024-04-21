import styled from 'styled-components';

export const ModalContainer = styled.div`
  position: absolute;
  top: 0;
  z-index: 10;
  height: 100%;
  width: 100%;
`;

export const ModalUnderlay = styled.section`
  background-color: #80808063;
  opacity: 90%;
  width: 100%;
  height: 100%;
`;

export const ModalCard = styled.div`
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
  border-radius: 25px;
  background-color: white;
  width: 150px;
  height: 150px;
  display: flex;
  justify-content: center;
  align-items: center;
`;
