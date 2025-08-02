import {StatusBar} from 'expo-status-bar';
import {Button, StyleSheet, Text, View} from 'react-native';
import React from 'react';
import {useEffect} from 'react';

import {ProgramList} from './features/program/Program';
import {Log} from './features/log/Log';
import {Stats} from './features/stats/Stats';
import {NavigationBar, Page} from './features/navigation-bar/NavigationBar';
import {theme} from './theme';
import {setSessionStorage, getSessionStorage} from './lib/session';


function CurrentPage({page}: {page: Page}): React.JSX.Element {
  switch (page) {
    case Page.Program:
      return <ProgramList />;
    case Page.Log:
      return <Log />;
    case Page.Stats:
      return <Stats />;
    default:
      return <Log />;
  }
}

export default function App() {
  const [page, setPage] = React.useState<Page>(() => {
    // Initialize page from session storage or default to Log
    if (typeof window !== 'undefined') {
      const storedPage = getSessionStorage<Page>('page');
      return storedPage || Page.Log;
    }
    return Page.Log;
  });

  if (typeof window !== 'undefined') {
    useEffect(() => {
      // Save the page to session storage whenever it changes
      setSessionStorage('page', page);
    }, [page]);
  }

  return (
    <View style={theme.container}>
      <StatusBar style="auto" />
      <CurrentPage page={page} />
      <NavigationBar page={page} setPage={setPage} />
    </View >
  );
}

