import 'package:flutter/material.dart';

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  const CustomAppBar ({super.key});

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight + 4);

  @override
  Widget build(BuildContext context) {
    return AppBar(
    backgroundColor: Color.fromARGB(255, 27, 27, 27),
      title: Row(
        children: [
          const Padding(
            padding: EdgeInsets.only(left: 8.0),
            child: Text(
              'Сервис авто "Третьяков"',
              style: TextStyle(
                fontSize: 28,
                color: Colors.white,
                fontWeight: FontWeight.bold,
              )
            ),
          ),
          const SizedBox(width: 26),
          TextButton(
            onPressed: () {
              // TODO: Do func
            },
            style: ButtonStyle(
              foregroundColor: WidgetStateProperty.resolveWith<Color>(
                (Set<WidgetState> states) {
                  if (states.contains(WidgetState.hovered)) {
                    return Colors.green;
                  }
                return Colors.white;
              },),
              overlayColor: WidgetStateProperty.all(Colors.transparent)
            ),
            child: const Text(
              'Каталог',
              style: TextStyle(fontSize: 18),
            ),
          ),
          const SizedBox(width: 26),
          TextButton(
            onPressed: () {
              // TODO:
            },
            style: ButtonStyle(
              foregroundColor: WidgetStateProperty.resolveWith<Color>(
                (Set<WidgetState> states) {
                  if (states.contains(WidgetState.hovered)) {
                    return Colors.green;
                  }
                  return Colors.white;
                }
              ),
              overlayColor: WidgetStateProperty.all(Colors.transparent)
            ),
            child: const Text(
              'О нас',
              style: TextStyle(fontSize: 18),
            )
          ),
          const Spacer(),
          const Text(
            '+7 908 254 31 81',
            style: TextStyle(fontSize: 18, color: Colors.white )
          )
        ],
      ),
      centerTitle: false,
      bottom: PreferredSize(
        preferredSize: const Size.fromHeight(4.0),
        child: Column(
          children: [
            Container(
              height: 4.0,
              color: Colors.green,
            ),
            Container(
              height: 4.0,
              color: Color.fromARGB(255, 27, 27, 27),
            )
          ],
        ),
      ),
    );
  }
}