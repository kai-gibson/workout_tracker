import React, {useState, useEffect} from 'react';
import {Button, Text, View, StyleSheet, Pressable, Dimensions} from 'react-native';
import {theme} from '../../theme';

function Program({name}: {name: string}): React.JSX.Element {
  return (
    <View style={theme.box}>
      <Text>{name}</Text>
    </View>
  );
}

function ProgramList() {
  const testProgramList = [
    "Upper Lower 4 Day Split",
    "Push Pull Legs 6 Day Split",
    "Full Body 3 Day Split",
  ]

  return (
    <View style={theme.container}>
      <Text style={theme.header}>Programs</Text>

      {testProgramList.map((name, index) => (
        <Program key={index} name={name} />
      ))}

      <Pressable style={theme.button} onPress={() => alert('Add Program')}>
        <Text>+</Text>
      </Pressable>
    </View >
  );
}

export {ProgramList, Program};
