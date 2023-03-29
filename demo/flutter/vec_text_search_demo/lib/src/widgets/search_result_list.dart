import 'package:flutter/material.dart';
import 'package:flutter/animation.dart';

class SearchResultList extends StatefulWidget {
  final List<Map<String, dynamic>> searchResults;

  SearchResultList({required this.searchResults});

  @override
  _SearchResultListState createState() => _SearchResultListState();
}

class _SearchResultListState extends State<SearchResultList>
    with TickerProviderStateMixin {
  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      shrinkWrap: true,
      itemCount: widget.searchResults.length,
      itemBuilder: (BuildContext context, int index) {
        final result = widget.searchResults[index];
        return Card(
          margin: const EdgeInsets.all(8.0),
          child: ListTile(
            title: Text(result['content']),
            subtitle: Text(
                'Certainty: ${(result['certainty'] * 100).toStringAsFixed(2)}%'),
            trailing:
                Text('Distance: ${result['distance'].toStringAsFixed(5)}'),
          ),
        );
      },
    );
  }
}
