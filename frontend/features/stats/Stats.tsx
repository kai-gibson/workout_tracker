import {View, Text} from 'react-native';
import {theme} from '../../theme';

function Stats() {
  return (
    <View style={theme.container}>
      <Text style={theme.header}>Stats</Text>
    </View>
  );
}

export {Stats};
