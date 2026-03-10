import React, { useState } from 'react';
import { View, Text, Button, Alert } from 'react-native';
import { BarCodeScanner } from 'expo-barcode-scanner';
import { submitTransaction } from '../services/blockchainService';

export default function QRScannerScreen({ navigation }) {
  const [hasPermission, setHasPermission] = useState(null);
  const [scanned, setScanned] = useState(false);

  const askForCameraPermission = () => {
    (async () => {
      const { status } = await BarCodeScanner.requestPermissionsAsync();
      setHasPermission(status === 'granted');
    })();
  };

  const handleBarCodeScanned = async ({ type, data }) => {
    setScanned(true);
    try {
      await submitTransaction(data);
      Alert.alert('Success', 'Carbon data logged successfully!');
      navigation.navigate('RouteOptimization');
    } catch (error) {
      Alert.alert('Error', 'Failed to log carbon data.');
    }
  };

  if (hasPermission === null) {
    return <View><Text>Requesting for camera permission</Text></View>;
  }
  if (hasPermission === false) {
    return <View><Text>No access to camera</Text></View>;
  }

  return (
    <View style={{ flex: 1 }}>
      <BarCodeScanner
        onBarCodeScanned={scanned ? undefined : handleBarCodeScanned}
        style={{ flex: 1 }}
      />
      {scanned && <Button title={'Tap to Scan Again'} onPress={() => setScanned(false)} />}
    </View>
  );
}