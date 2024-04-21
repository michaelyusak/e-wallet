import React, { useEffect, useState } from 'react';

import showNavIcon from '../../assets/img/show-nav.png';
import hideNavIcon from '../../assets/img/hide-nav.png';
import { INavItem } from '../../interfaces/NavItem';
import * as S from './SideNav.Styles';
import { NavLink } from 'react-router-dom';

type SideNavProps = {
  navItems: INavItem[];
};

const SideNav = ({ navItems }: SideNavProps): React.ReactElement => {
  const [isExpanded, setExpanded] = useState<boolean>(false);

  function useWindowDimensions() {
    const [windowDimension, setWindowDimensions] = useState(window.innerWidth);

    useEffect(() => {
      function handleResize() {
        setWindowDimensions(window.innerWidth);
      }

      window.addEventListener('resize', handleResize);
      return () => window.removeEventListener('resize', handleResize);
    }, []);

    return windowDimension;
  }

  function handleSetExpanded(): void {
    setExpanded(!isExpanded);
  }

  const windowWidth = useWindowDimensions();

  return (
    <S.SideNavContainer>
      {windowWidth > 710 ? (
        <S.SideNavigation>
          <h1>Sea Wallet</h1>

          <div>
            {navItems.map((navItem) => (
              <NavLink
                key={navItem.itemName}
                end={navItem.isDefault && true}
                to={navItem.path}
                onClick={navItem.onClick}
                className={({ isActive }) => (isActive ? 'isOn' : 'isOff')}
              >
                <img src={navItem.itemIconSrc} alt=""></img>

                <p>{navItem.itemName}</p>
              </NavLink>
            ))}
          </div>
        </S.SideNavigation>
      ) : !isExpanded ? (
        <S.NarrowedMobileNav>
          <h1>Sea Wallet</h1>

          <div>
            {navItems.map((navItem) => (
              <NavLink
                key={navItem.itemName}
                end={navItem.isDefault && true}
                to={navItem.path}
                onClick={navItem.onClick}
                className={({ isActive }) => (isActive ? 'isOn' : 'isOff')}
              >
                <img src={navItem.itemIconSrc} alt=""></img>
              </NavLink>
            ))}
          </div>

          <button onClick={() => handleSetExpanded()}>
            <img src={isExpanded ? hideNavIcon : showNavIcon} alt=""></img>
          </button>
        </S.NarrowedMobileNav>
      ) : (
        <S.ExpandedMobileNav>
          <h1>Sea Wallet</h1>

          <div>
            {navItems.map((navItem) => (
              <NavLink
                key={navItem.itemName}
                end={navItem.isDefault && true}
                to={navItem.path}
                className={({ isActive }) => (isActive ? 'isOn' : 'isOff')}
              >
                <img src={navItem.itemIconSrc} alt=""></img>

                <p>{navItem.itemName}</p>
              </NavLink>
            ))}
          </div>

          <button onClick={() => handleSetExpanded()}>
            <img src={isExpanded ? hideNavIcon : showNavIcon} alt=""></img>
          </button>
        </S.ExpandedMobileNav>
      )}
    </S.SideNavContainer>
  );
};

export default SideNav;
