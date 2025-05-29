import 'package:flutter/material.dart';

class CarAdvertisement extends StatelessWidget {
  const CarAdvertisement({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox.expand(
      child: Row(
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

              child: Stack(
                children: [
                  Transform.translate(
                    offset: const Offset(-40, 0),
                    child: Image.asset(
                      'assets/CarRight.png',
                      width: double.infinity,
                      height: double.infinity,
                    ),
                  ),

                  Positioned(
                    left: 20,
                    bottom: 450,
                    child: Text(
                      'выбор',
                      style: TextStyle(
                        color: Colors.green,
                        fontSize: 48
                      ),
                    ),
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