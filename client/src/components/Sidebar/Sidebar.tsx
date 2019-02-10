import React from 'react';
import { Book, Cafeteria } from 'grommet-icons';
import { SidebarWrapper, SidebarLink, SidebarLogo } from './sidebar.styles';

const Sidebar: React.FC = () => (
  <SidebarWrapper>
    <SidebarLogo>
      <Cafeteria color="light-1" size="medium" style={{ margin: '0 5px' }} />
      <Book color="light-1" size="medium" style={{ margin: '0 5px' }} />
    </SidebarLogo>
    <div>
      <SidebarLink to="/">Home</SidebarLink>
    </div>
    <div>
      <SidebarLink to="/recipes">Recipes</SidebarLink>
    </div>
    <div>
      <SidebarLink to="/ingredients">Ingredients</SidebarLink>
    </div>
  </SidebarWrapper>
);

export default Sidebar;
