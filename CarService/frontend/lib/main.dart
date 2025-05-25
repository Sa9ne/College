import 'package:flutter/material.dart';
import 'package:frontend/Screens/Car_Service_Screen.dart';

void main() {
  runApp(const CarService());
}

class CarService extends StatelessWidget {
  const CarService({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Car Service',
      theme: ThemeData(
        scaffoldBackgroundColor: const Color.fromARGB(255, 27, 27, 27),
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.green),
        useMaterial3: true,
      ),
      home: CarServiceScreen()
    );
  }
}

