import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import QRScannerScreen from './screens/QRScannerScreen';

const Stack = createStackNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="QRScanner">
        <Stack.Screen name="QRScanner" component={QRScannerScreen} options={{ title: 'Scan Shipment QR' }} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}