import 'package:flutter/material.dart';

class CarAdvertisement extends StatelessWidget {
  const CarAdvertisement({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox.expand(
      child: Stack(
        children: [
          Row(
            children: [
              Expanded(
                flex: 1,
                child: Container(
                  color: Colors.green
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  color: Color.fromARGB(255, 27, 27, 27),
                )
              ),
            ]
          ),
          
          Positioned.fill(
            child: Align(
              alignment: Alignment.center,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  ShaderMask(
                    shaderCallback: (Rect bounds) {
                      return LinearGradient(
                        colors: [Colors.black, Colors.green],
                        stops: [0.5, 0.5],
                      ).createShader(bounds);
                    },
                    blendMode: BlendMode.srcIn,
                    child: const Text(
                      'Большой выбор\n        На любой вкус',
                      style: TextStyle(
                        fontSize: 58,
                        color: Colors.white,
                        fontWeight: FontWeight.bold
                      ),
                    ),
                  ),

                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Image.asset(
                        'assets/CarLeft.png',
                        height: MediaQuery.of(context).size.height * 0.6,
                        width: MediaQuery.of(context).size.width * 0.3,
                      ),

                      Image.asset(
                        'assets/CarRight.png',
                        height: MediaQuery.of(context).size.height * 0.6,
                        width: MediaQuery.of(context).size.width * 0.3,
                      )
                    ],
                  )
                ],
              )
            ),
          )
        ],
      ),
    );
  }
}