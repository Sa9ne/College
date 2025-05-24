import 'package:flutter/material.dart';

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
      home: const HigherMenu(),
    );
  }
}

class HigherMenu extends StatelessWidget {
  const HigherMenu({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.green,
        title: const Padding(
          padding: EdgeInsets.only(left: 8.0),
          child: Text(
            'Сервис Подержаных авто',
            style: TextStyle(
              fontSize: 28,
              color: Colors.black,
              fontWeight: FontWeight.bold,
            )
          ),
        ),
        centerTitle: false,
      )
    );
  }
}