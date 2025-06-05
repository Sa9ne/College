import 'package:flutter/material.dart';
import 'package:frontend/services/api_service.dart';

Future<void> showMyDialog(BuildContext context) async {
  final orders = await ShowBoughtCar();

  showDialog(
    context: context,
    barrierDismissible: true,
    builder: (BuildContext context) {
      return AlertDialog(
        content: SizedBox(
          width: double.maxFinite,
          child: Stack(
            children: [
              Padding(
                padding: const EdgeInsets.only(top: 40.0),
                child: ListView(
                  shrinkWrap: true,
                  children: [
                    const Center(
                      child: Text(
                        "Все заказы",
                        style: TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),
                    ...orders.map((order) => Card(
                      elevation: 4,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.all(16.0),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              "Номер заказа: ${order.id}",
                              style: const TextStyle(
                                  fontSize: 16,
                                  fontWeight: FontWeight.w500),
                            ),
                            const SizedBox(height: 8),
                            Text("Бренд: ${order.brand}", style: const TextStyle(fontSize: 16)),
                            Text("Модель: ${order.name}", style: const TextStyle(fontSize: 16)),
                            const SizedBox(height: 8),
                            Text(
                              "Номер клиента: ${order.phone}",
                              style: const TextStyle(fontSize: 16, color: Colors.grey),
                            ),
                          ],
                        ),
                      ),
                    )).toList(),
                  ],
                ),
              ),
              Positioned(
                top: 0,
                right: 0,
                child: IconButton(
                  icon: const Icon(Icons.close),
                  onPressed: () => Navigator.of(context).pop(),
                ),
              ),
            ],
          ),
        ),
      );
    },
  );
}
