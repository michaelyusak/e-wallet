import styled from 'styled-components';

export const SideNavContainer = styled.div`
  position: fixed;
  top: 0;

  @media (max-width: 710px) {
    z-index: 2;
  }
`;

export const SideNavigation = styled.nav`
  width: 280px;
  padding: 30px 27px;
  border-radius: 0px 20px 10px 0px;
  background-color: #f6f4f5;
  height: 100vh;

  h1 {
    line-height: 42px;
    font-size: 28px;
    font-weight: 600;
    color: #4d47c3;
    width: 152px;
  }

  div {
    list-style: none;
    margin-top: 249px;
    display: flex;
    flex-direction: column;
    gap: 25px;

    a {
      display: flex;
      gap: 18px;
      padding: 4px 0px;
      text-decoration: none;

      img {
        width: 24px;
      }
    }

    a.isOn {
      img {
        filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
          saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      }

      p {
        color: #4d47c3;
        line-height: 27px;
        font-size: 18px;
        font-weight: 600;
      }
    }

    a.isOff {
      p {
        color: #95999e;
        line-height: 27px;
        font-size: 18px;
        font-weight: 600;
      }
    }

    a:hover {
      filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
        saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      cursor: pointer;
    }
  }
`;

export const NarrowedMobileNav = styled.nav`
  padding: 30px 25px;
  border-radius: 0px 20px 10px 0px;
  background-color: #f6f4f5;
  display: flex;
  height: 100vh;
  width: 74px;
  flex-direction: column;
  align-items: center;

  h1 {
    font-weight: 600;
    color: #4d47c3;
    width: 46px;
    line-height: 21px;
    font-size: 14px;
  }

  div {
    list-style: none;
    margin-top: 249px;
    width: fit-content;
    display: flex;
    flex-direction: column;
    gap: 25px;

    a {
      display: flex;
      gap: 18px;
      padding: 4px 0px;
      height: 27px;
      align-items: center;

      img {
        width: 24px;
      }
    }

    a.isOn {
      img {
        filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
          saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      }
    }

    a:hover {
      filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
        saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      cursor: pointer;
    }
  }

  button:nth-child(3) {
    margin-top: 31vh;
    border: none;
    outline: none;
    background-color: transparent;
    height: 27px;
    cursor: pointer;
  }
`;

export const ExpandedMobileNav = styled.nav`
  padding: 30px 25px;
  border-radius: 0px 20px 10px 0px;
  background-color: #f6f4f5;
  height: 100vh;
  width: 269px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;

  h1 {
    line-height: 42px;
    font-size: 28px;
    font-weight: 600;
    color: #4d47c3;
  }

  div {
    list-style: none;
    margin-top: 249px;
    width: fit-content;
    display: flex;
    flex-direction: column;
    gap: 25px;

    a {
      display: flex;
      gap: 18px;
      padding: 4px 0px;
      height: 27px;
      align-items: center;
      text-decoration: none;

      img {
        width: 24px;
      }

      p {
        color: #95999e;
        line-height: 27px;
        font-size: 18px;
        font-weight: 600;
      }
    }

    a.isOn {
      img {
        filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
          saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      }

      p {
        color: #4d47c3;
        line-height: 27px;
        font-size: 18px;
        font-weight: 600;
      }
    }

    a.isOff {
      p {
        color: #95999e;
        line-height: 27px;
        font-size: 18px;
        font-weight: 600;
      }
    }

    a:hover {
      filter: brightness(0) saturate(100%) invert(20%) sepia(100%)
        saturate(2440%) hue-rotate(238deg) brightness(85%) contrast(78%);
      cursor: pointer;
    }
  }

  button:nth-child(3) {
    margin-top: 300px;
    border: none;
    outline: none;
    background-color: transparent;
    height: 27px;
    cursor: pointer;
  }
`;
