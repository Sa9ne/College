import 'package:flutter/material.dart';
import 'package:frontend/models/car.dart';
import 'package:frontend/services/api_service.dart';
import 'package:frontend/widgets/Client.dart';

class Catalog extends StatefulWidget {
  const Catalog({super.key});

  @override
  State<Catalog> createState() => _CarCatalogState();
}

class _CarCatalogState extends State<Catalog> {
  late Future<List<Car>> futureCars;

  @override
  void initState() {
    super.initState();
    futureCars = fetchCars();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<Car>>(
      future: futureCars,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Padding(
            padding: EdgeInsets.all(20),
            child: CircularProgressIndicator(),
          );
        } else if (snapshot.hasError) {
          return Text("Ошибка загрузки: ${snapshot.error}");
        } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
          return const Text("Нет доступных машин");
        }

        final cars = snapshot.data!;

        return Container(
          color: Color.fromARGB(255, 27, 27, 27),
          padding: const EdgeInsets.symmetric(vertical: 20),
          child: Column(
            children: cars.map((car) => CarCard(car: car)).toList(),
          ),
        );
      },
    );
  }
}

class CarCard extends StatelessWidget {
  final Car car;

  const CarCard({super.key, required this.car});

  @override
  Widget build(BuildContext context) {
    return Card(
      color: Colors.white,
      margin: const EdgeInsets.all(10),
      elevation: 3,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(12),
      ),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    "${car.brand} ${car.name}",
                    style: const TextStyle(fontSize: 22, fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 16),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            _iconText(Icons.speed, "${car.power} л.с."),
                            const SizedBox(height: 10),
                            _iconText(Icons.color_lens, car.color),
                            const SizedBox(height: 10),
                            _iconText(Icons.directions_car, "${car.mileage} км"),
                          ],
                        ),
                      ),
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            _iconText(Icons.calendar_today, "${car.year} г."),
                            const SizedBox(height: 10),
                            _iconText(Icons.person, "${car.owners} владелец(а)"),
                            const SizedBox(height: 10),
                            _iconText(Icons.info, car.condition),
                          ],
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: 16),
                  Text(
                    "Цена: ${car.price} ₽",
                    style: const TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
                  ),
                ],
              ),
            ),

            const SizedBox(width: 20),

            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                SizedBox(
                  width: 140,
                  height: 50,
                  child: OutlinedButton(
                    onPressed: () => showClientFormDialog(context, car.vin),
                    style: OutlinedButton.styleFrom(
                      side: const BorderSide(color: Colors.green, width: 2),
                      foregroundColor: Colors.black,
                      padding: const EdgeInsets.symmetric(horizontal: 16),
                      textStyle: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(10),
                      ),
                    ),
                    child: const Text("Заказать"),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  Widget _iconText(IconData icon, String text) {
    return Row(
      children: [
        Icon(icon, size: 22, color: Colors.grey[700]),
        const SizedBox(width: 8),
        Expanded(
          child: Text(
            text,
            style: const TextStyle(fontSize: 16),
            overflow: TextOverflow.ellipsis,
          ),
        ),
      ],
    );
  }
}
