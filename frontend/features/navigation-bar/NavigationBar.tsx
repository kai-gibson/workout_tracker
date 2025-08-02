import React from 'react';
import {View, Text, Pressable} from 'react-native';

import {theme} from '../../theme';

enum Page {
  Program = 'Program',
  Log = 'Log',
  Stats = 'Stats',
  Settings = 'Settings',
}

const NAVIGATION_BUTTONS: Page[] = [
  Page.Program,
  Page.Log,
  Page.Stats,
]

function NavigationButton(
  {selectedPage, page, setPage}: {selectedPage: Page, page: Page, setPage: (page: Page) => void}
): React.JSX.Element {
  const [textStyle, buttonStyle] = page == selectedPage
    ? [theme.selectedNavigationText, theme.selectedNavigationButton]
    : [theme.navigationText, theme.navigationButton];

  return (
    <Pressable style={buttonStyle} onPress={() => setPage(page)}>
      <Text style={textStyle}>{page}</Text>
    </Pressable>
  );
}

function NavigationBar(
  {page, setPage}: {page: Page, setPage: (page: Page) => void}
): React.JSX.Element {
  return (
    <View style={theme.navigationBar} >
      {NAVIGATION_BUTTONS.map((button) => {
        return <NavigationButton selectedPage={page} page={button} setPage={setPage} />
      })}
    </View >
  );
}

export {NavigationBar, Page};
