import 'dart:convert';
import 'package:frontend/models/car.dart';
import 'package:http/http.dart' as http;

Future<List<Car>> fetchCars() async {
  final response = await http.get(Uri.parse('http://localhost:8080/Catalog'));

  if (response.statusCode == 200) {
    final List<dynamic> jsonData = json.decode(response.body);
    return jsonData.map((carJson) => Car.fromJson(carJson)).toList();
  } else {
    throw Exception('Failed to load cars');
  }
}