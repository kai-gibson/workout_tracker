import {StyleSheet, Dimensions} from 'react-native';

const screenWidth = Dimensions.get('window').width;

const theme = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    justifyContent: 'center',
    alignItems: 'center',
  },
  "box": {
    borderWidth: 1,
    borderColor: '#000',
    borderRadius: 5,
    padding: 10,
    width: screenWidth * 0.8,
    margin: 5,
    alignItems: 'center',
  },
  "header": {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 10,
    alignSelf: 'center',
    position: 'absolute',
    top: 20,
  },
  "button": {
    borderWidth: 1,
    borderColor: '#000',
    borderRadius: 5,
    padding: 10,
    margin: 5,
    width: screenWidth * 0.04,
    alignItems: 'center',
  },
  "navigationBar": {
    flexDirection: 'row',
    justifyContent: 'space-around',
    borderTopWidth: 1,
    borderColor: '#000',
    width: screenWidth,
    position: 'absolute',
    bottom: 0,
  },
  "navigationText": {
    fontSize: 16,
    fontWeight: 'bold',
  },
  "selectedNavigationText": {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#fff',
  },
  "navigationButton": {
    width: screenWidth * (1 / 3),
    height: screenWidth * 0.05,
    justifyContent: 'center',
    alignItems: 'center',
  },
  "selectedNavigationButton": {
    width: screenWidth * (1 / 3),
    height: screenWidth * 0.05,
    justifyContent: 'center',
    alignItems: 'center',
    color: '#000',
    backgroundColor: '#000',
  }
});

export {theme, screenWidth};
