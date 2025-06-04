import "package:flutter/material.dart";
import "package:frontend/widgets/AboutUs.dart";
import "package:frontend/widgets/CustomMainCont.dart";
import "package:frontend/widgets/CustomAppBar.dart";
import 'package:frontend/widgets/Catalog.dart';

class CarServiceScreen extends StatelessWidget {
  CarServiceScreen({super.key});

  final ScrollController _scrollController = ScrollController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(scrollController: _scrollController,),
      body: SingleChildScrollView(
        controller: _scrollController,
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