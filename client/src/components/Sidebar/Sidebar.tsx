import React from 'react';
import { Book, Cafeteria } from 'grommet-icons';
import { SidebarWrapper } from './sidebar.styles';

const Sidebar: React.FC = () => (
  <SidebarWrapper>
    <div>
      <Cafeteria color="light-1" size="medium" style={{ margin: '0 5px' }} />
      <Book color="light-1" size="medium" style={{ margin: '0 5px' }} />
    </div>
  </SidebarWrapper>
);

export default Sidebar;
