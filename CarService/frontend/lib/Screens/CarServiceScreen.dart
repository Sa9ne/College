import "package:flutter/material.dart";
import "package:frontend/widgets/AboutUs.dart";
import "package:frontend/widgets/CustomMainCont.dart";
import "package:frontend/widgets/CustomAppBar.dart";
import 'package:frontend/widgets/Catalog.dart';

class CarServiceScreen extends StatelessWidget {
  const CarServiceScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: SingleChildScrollView(
        child: Column(
          children: [
            CarAdvertisement(),
            Catalog(),
            AboutUs(),
          ],
        ),
      )
    );
  }
}