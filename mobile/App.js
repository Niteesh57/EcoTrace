import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import QRScannerScreen from './screens/QRScannerScreen';
import RouteOptimizationScreen from './screens/RouteOptimizationScreen';

const Stack = createStackNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="QRScanner">
        <Stack.Screen name="QRScanner" component={QRScannerScreen} />
        <Stack.Screen name="RouteOptimization" component={RouteOptimizationScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}