import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Grommet, Box } from 'grommet';
import { grommet } from 'grommet/themes';
import { routes, IRouteConfig } from './routes';
import Sidebar from './components/Sidebar/Sidebar';
import { ContentWrapper } from './app.styles';

const App: React.FC = () => (
  <Router>
    <Grommet full theme={grommet}>
      <Box direction="row" align="stretch" style={{ height: '100%' }}>
        <Box basis="small" align="stretch">
          <Sidebar />
        </Box>
        <Box>
          <ContentWrapper>
            {routes.map((route: IRouteConfig, i: number) => (
              <Route
                key={i}
                path={route.path}
                exact={route.exact}
                component={route.component}
              />
            ))}
          </ContentWrapper>
        </Box>
      </Box>
    </Grommet>
  </Router>
);

export default App;
