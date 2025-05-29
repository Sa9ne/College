import "package:flutter/material.dart";
import "package:frontend/widgets/CustomMainCont.dart";
import "package:frontend/widgets/Custom_app_bar.dart";

class CarServiceScreen extends StatelessWidget {
  const CarServiceScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: const CarAdvertisement()
    );
  }
}