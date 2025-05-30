import 'package:flutter/material.dart';

class AboutUs extends StatelessWidget {
  const AboutUs({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Color.fromARGB(255, 27, 27, 27),
      width: double.infinity,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: const [
          Divider(
            color: Colors.green,
            thickness: 2,
            height: 20,
          ),
          Text(
            'О нас',
            style: TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.bold,
              color: Colors.white,
            ),
          ),
          SizedBox(height: 12),
          Text(
            'Мы — команда, которая предоставляет качественные автомобили и отличный сервис.\n'
            'С 2010 года мы помогаем клиентам найти идеальное авто.',
            style: TextStyle(
              fontSize: 16,
              color: Colors.white70,
            ),
          ),
          SizedBox(height: 20),
          Text(
            'Контакты:\n📧 trealev@gmail.com\n📞 +7 (908) 254-31-81',
            style: TextStyle(
              fontSize: 16,
              color: Colors.white70,
            ),
          ),
        ],
      ),
    );
  }
}