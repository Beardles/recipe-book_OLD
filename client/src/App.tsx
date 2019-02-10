import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Grommet, Grid, Box } from 'grommet';
import { grommet } from 'grommet/themes';
import { routes, IRouteConfig } from './routes';
import Sidebar from './components/Sidebar/Sidebar';
import { ContentWrapper } from './app.styles';

const App: React.FC = () => (
  <Router>
    <Grommet full theme={grommet}>
      <Grid
        fill
        rows={['auto', 'flex']}
        columns={['auto', 'flex']}
        areas={[
          { name: 'sidebar', start: [0, 1], end: [0, 1] },
          { name: 'main', start: [1, 1], end: [1, 1] },
        ]}
      >
        <Box gridArea="sidebar" width="small">
          <Sidebar />
        </Box>
        <Box gridArea="main">
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
      </Grid>
    </Grommet>
  </Router>
);

export default App;
