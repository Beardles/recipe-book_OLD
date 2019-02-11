import styled from 'styled-components';
import { Link } from 'react-router-dom';

export const SidebarWrapper = styled.div`
  background-color: #7d4cdb;
  height: 100%;
  padding: 50px 0;
  text-align: center;
`;

export const SidebarLogo = styled.div`
  margin-bottom: 25px;
`;

export const SidebarLink = styled(Link)`
  color: #ffffff;
  display: inline-block;
  padding: 10px 0;
  text-decoration: none;
  width: 100%;

  &:hover {
    background-color: #9766f5;
    cursor: pointer;
    text-decoration: none;
  }

  &:active {
    color: #7d4cdb;
  }
`;
