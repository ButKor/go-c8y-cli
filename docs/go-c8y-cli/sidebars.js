module.exports = {
  someSidebar: {
    Documentation: [
      'introduction',
      'demo',
      {
        type: 'category',
        label: 'Installation',
        collapsed: true,
        items: [
          {
            type: 'autogenerated',
            dirName: 'installation',
          },
        ],
      },
      'gettingstarted',
      {
        type: 'category',
        label: 'Concepts',
        collapsed: false,
        items: [
          {
            type: 'autogenerated',
            dirName: 'concepts', // Generate sidebar slice from docs/api/c8y
          },
        ],
      },
      {
        type: 'category',
        label: 'Configuration',
        collapsed: true,
        items: [
          {
            type: 'autogenerated',
            dirName: 'configuration', // Generate sidebar slice from docs/api/c8y
          },
        ],
      },
      {
        type: 'category',
        label: 'Troubleshooting',
        collapsed: true,
        items: [
          {
            type: 'autogenerated',
            dirName: 'troubleshooting', // Generate sidebar slice from docs/api/c8y
          },
        ],
      },
      {
        type: 'category',
        label: 'Examples',
        collapsed: true,
        items: [
          {
            type: 'autogenerated',
            dirName: 'examples',
          },
        ],
      },
      {
        type: 'category',
        label: 'Tutorials',
        collapsed: true,
        items: [
          {
            type: 'autogenerated',
            dirName: 'tutorials',
          },
        ],
      },
    ],
    // Commands: [
    //   {
    //     type: 'category',
    //     label: 'c8y (native)',
    //     collapsed: true,
    //     items: [
    //       {
    //         type: 'autogenerated',
    //         dirName: 'api/c8y', // Generate sidebar slice from docs/api/c8y
    //       },
    //     ],
    //   },
    // ],
  },
  cli: [
    'cli',
    {
      type: 'category',
      label: 'CLI',
      collapsed: false,
      items: [
        {
          type: 'autogenerated',
          dirName: 'cli',
        },
        // {
        //     type: 'autogenerated',
        //     dirName: 'cli/psc8y',
        //   },
      ],
    },
  ],
};
