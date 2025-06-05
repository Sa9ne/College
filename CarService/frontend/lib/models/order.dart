class Order {
  final int id;
  final String brand;
  final String vin;
  final String phone;
  final String name;

  Order({
    required this.id,
    required this.brand,
    required this.vin,
    required this.phone,
    required this.name,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      id: json['id'],
      brand: json['brand'],
      name: json['name'],
      vin: json['vin'],
      phone: json['phone'],
    );
  }
}
