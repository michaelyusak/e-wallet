import styled from 'styled-components';

export const FormFields = styled.form<{ $Gap?: string }>`
  display: flex;
  flex-direction: column;
  gap: ${(props) => props.$Gap || '38px'};
`;

export const Button = styled.input.attrs({ type: 'submit' })`
  border: none;
  outline: none;
  border-radius: 9px;
  height: 59px;
  background-color: #4d47c3;
  box-shadow: 0px 2px 40px 0px #4d47c366;

  color: white;
  line-height: 24px;
  font-size: 16px;
  font-weight: 500;

  cursor: pointer;
`;
