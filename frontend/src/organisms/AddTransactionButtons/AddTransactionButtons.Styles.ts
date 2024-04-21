import styled from 'styled-components';

export const Button = styled.button`
  padding: 10px;
  background-color: #4d47c3;
  border: solid #95999e 0.5px;
  width: 140px;
  height: 44px;
  display: flex;
  gap: 10px;
  border-radius: 15px;
  align-items: center;
  cursor: pointer;
  outline: none;

  img {
    width: 26px;
    filter: brightness(0) saturate(100%) invert(100%) sepia(0%) saturate(7130%)
      hue-rotate(87deg) brightness(106%) contrast(99%);
  }

  p {
    line-height: 24px;
    font-size: 16px;
    font-weight: 500;
    color: white;
  }
`;
