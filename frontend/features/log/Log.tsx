import {View, Text} from 'react-native';
import {theme} from '../../theme';

function Log() {
  return (
    <View style={theme.container}>
      <Text style={theme.header}>Log</Text>
    </View>
  );
}

export {Log};
