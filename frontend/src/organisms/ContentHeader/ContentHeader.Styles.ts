import styled from 'styled-components';

export const ContentHeader = styled.header`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-left: 280px;
  padding: 30px 60px 20px 65px;

  h2 {
    color: #95999e;
    line-height: 36px;
    font-size: 24px;
    font-weight: 600;
    margin: auto 0;
  }

  @media (max-width: 710px) {
    margin-bottom: 35px;

    h2 {
      line-height: 30px;
      font-size: 20px;
    }
  }
`;

export const ProfileNav = styled.div`
  height: 45px;
  display: flex;
  flex-direction: column;
  gap: 5px;
  align-items: flex-end;

  > button {
    border: none;
    outline: none;
    background-color: transparent;
    height: 45px;
    border-radius: 100%;

    > img {
      width: 45px;
      height: 45px;
      object-fit: cover;
      border-radius: 100%;
      cursor: pointer;
    }
  }

  ul {
    list-style: none;
    background-color: white;
    border-radius: 0px 0px 10px 10px;
    width: 132px;
    z-index: 3;
    padding: 6px 20px;
    display: flex;
    flex-direction: column;
    gap: 18px;
    border: 0.5px solid #95999e;

    li {
      button {
        display: flex;
        justify-content: space-between;
        align-items: center;
        height: fit-content;
        width: 100%;
        cursor: pointer;
        border: none;
        outline: none;
        background-color: transparent;
      }

      p {
        color: #282828;
        line-height: 22.5px;
        font-size: 15px;
        font-weight: 600;
      }
    }

    li:nth-child(1) {
      img {
        width: 17px;
        height: auto;
      }
    }

    li:nth-child(2) {
      img {
        width: 21.59px;
        height: auto;
      }
    }
  }
`;
